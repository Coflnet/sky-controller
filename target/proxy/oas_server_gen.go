// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"time"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APITimeGet implements GET /Api/time operation.
	//
	// GET /Api/time
	APITimeGet(ctx context.Context) (time.Time, error)
	// BaseAhPlayerIdPost implements POST /Base/ah/{playerId} operation.
	//
	// Request ah update for a player.
	//
	// POST /Base/ah/{playerId}
	BaseAhPlayerIdPost(ctx context.Context, params BaseAhPlayerIdPostParams) error
	// BaseKeyPost implements POST /Base/key operation.
	//
	// POST /Base/key
	BaseKeyPost(ctx context.Context, req *KeyCreate) error
	// BaseKeysPartyCountGet implements GET /Base/keys/{party}/count operation.
	//
	// GET /Base/keys/{party}/count
	BaseKeysPartyCountGet(ctx context.Context, params BaseKeysPartyCountGetParams) (int32, error)
	// ProxyHypixelAhPlayerPlayerUuidGet implements GET /Proxy/hypixel/ah/player/{playerUuid} operation.
	//
	// Load the auctions of a player.
	//
	// GET /Proxy/hypixel/ah/player/{playerUuid}
	ProxyHypixelAhPlayerPlayerUuidGet(ctx context.Context, params ProxyHypixelAhPlayerPlayerUuidGetParams) ([]SaveAuction, error)
	// ProxyHypixelGet implements GET /Proxy/hypixel operation.
	//
	// Proxy the path to hypixel using the assigned key (no key required).
	//
	// GET /Proxy/hypixel
	ProxyHypixelGet(ctx context.Context, params ProxyHypixelGetParams) (ProxyHypixelGetRes, error)
	// ProxyHypixelStatusGet implements GET /Proxy/hypixel/status operation.
	//
	// Can proxy anything.
	//
	// GET /Proxy/hypixel/status
	ProxyHypixelStatusGet(ctx context.Context) error
	// SkyblockAuctionsGet implements GET /skyblock/auctions operation.
	//
	// GET /skyblock/auctions
	SkyblockAuctionsGet(ctx context.Context) (SkyblockAuctionsGetRes, error)
	// SkyblockAuctionsPost implements POST /skyblock/auctions operation.
	//
	// POST /skyblock/auctions
	SkyblockAuctionsPost(ctx context.Context) (SkyblockAuctionsPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}