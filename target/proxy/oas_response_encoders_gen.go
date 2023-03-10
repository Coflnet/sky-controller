// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/json"
)

func encodeAPITimeGetResponse(response time.Time, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	json.EncodeDateTime(e, response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeBaseAhPlayerIdPostResponse(response *BaseAhPlayerIdPostOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeBaseKeyPostResponse(response *BaseKeyPostOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeBaseKeysPartyCountGetResponse(response int32, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	e.Int32(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeProxyHypixelAhPlayerPlayerUuidGetResponse(response []SaveAuction, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	e.ArrStart()
	for _, elem := range response {
		elem.Encode(e)
	}
	e.ArrEnd()
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeProxyHypixelGetResponse(response ProxyHypixelGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ProxyHypixelGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ProxyHypixelGetOKTextJSON:
		w.Header().Set("Content-Type", "text/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ProxyHypixelGetOKTextPlain:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeProxyHypixelStatusGetResponse(response *ProxyHypixelStatusGetOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeSkyblockAuctionsGetResponse(response SkyblockAuctionsGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SkyblockAuctionsGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SkyblockAuctionsGetOKTextJSON:
		w.Header().Set("Content-Type", "text/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SkyblockAuctionsGetOKTextPlain:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSkyblockAuctionsPostResponse(response SkyblockAuctionsPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SkyblockAuctionsPostOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SkyblockAuctionsPostOKTextJSON:
		w.Header().Set("Content-Type", "text/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SkyblockAuctionsPostOKTextPlain:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
