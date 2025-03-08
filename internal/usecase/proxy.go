package usecase

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Coflnet/sky-controller/internal/metrics"
	"github.com/Coflnet/sky-controller/internal/utils"
	api "github.com/Coflnet/sky-controller/target/proxy"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ProxyScaler struct {
	Interval        time.Duration
	DesiredReplicas int
	CurrentReplicas int
}

var (
	proxyClient *api.Client
	clientset   *kubernetes.Clientset
)

func (p *ProxyScaler) Start() {

	p.init()

	go func() {
		for {
			err := p.scaleIfNecessary()

			if err != nil {
				log.Error().Err(err).Msgf("Error while scaling proxy")
			}

			time.Sleep(p.Interval)
		}
	}()
}

func (p *ProxyScaler) init() {

	// init api client
	var err error
	proxyClient, err = api.NewClient(utils.ProxyBaseURL())
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create proxy client")
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create in-cluster config")
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create kubernetes client")
	}
}

func (p *ProxyScaler) scaleIfNecessary() error {

	// update desired and current replica count
	err := p.updateReplicaCounts()
	if err != nil {
		return err
	}

	// update if necessary
	if p.DesiredReplicas != p.CurrentReplicas {
		log.Info().Msgf("Scaling proxy from %d to %d replicas", p.CurrentReplicas, p.DesiredReplicas)
		err = p.scale()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProxyScaler) scale() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get deployment
	deployment, err := clientset.AppsV1().Deployments("sky").Get(ctx, "sky-proxy", metav1.GetOptions{})

	if err != nil {
		return err
	}

	// scale deployment
	deployment.Spec.Replicas = int32Ptr(int32(p.DesiredReplicas))
	_, err = clientset.AppsV1().Deployments("sky").Update(ctx, deployment, metav1.UpdateOptions{})

	if err != nil {
		log.Error().Err(err).Msg("Failed to update proxy deployment")
		return err
	}

	return nil
}

func (p *ProxyScaler) updateReplicaCounts() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// execute both requests in parallel
	// and return for an error if one occurs

	errChan := make(chan error, 2)

	go func(c context.Context) {
		err := p.updateCurrentReplicaCount(c)
		errChan <- err
	}(ctx)

	go func(c context.Context) {
		err := p.updateDesiredReplicaCount(c)
		errChan <- err
	}(ctx)

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProxyScaler) updateDesiredReplicaCount(ctx context.Context) error {
	proxyCount, err := proxyClient.BaseKeysPartyCountGet(ctx, api.BaseKeysPartyCountGetParams{
		Party: "hypixel",
	})

	if err != nil {
		return err
	}
	log.Info().Msgf("Found %d keys for party hypixel", proxyCount)

	nodeCount, err := p.listNodeCount(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to list nodes, using default value")
		nodeCount = 6
	}

	// 3 masters, so -3
	nodeCount -= 3

	count := int(proxyCount)
	if nodeCount < count {
		count = nodeCount
	}

	log.Info().Msgf("Found %d keys for party hypixel", count)
	p.DesiredReplicas = int(count)
	return nil
}

func (p *ProxyScaler) listNodeCount(ctx context.Context) (int, error) {
	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})

	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return 0, err
	}

	log.Info().Msgf("Found %d nodes", len(nodes.Items))
	return len(nodes.Items), nil
}

func (p *ProxyScaler) updateCurrentReplicaCount(ctx context.Context) error {
	// get deployment
	deployment, err := clientset.AppsV1().Deployments("sky").Get(ctx, "sky-proxy", metav1.GetOptions{})

	if err != nil {
		log.Error().Err(err).Msg("Failed to get proxy deployment")
		return err
	}

	log.Debug().Msgf("Found %d replicas for proxy deployment", deployment.Status.Replicas)
	p.CurrentReplicas = int(deployment.Status.Replicas)

	metrics.UpdateProxyReplicas(p.CurrentReplicas)
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
