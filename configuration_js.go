// +build js,wasm

package webrtc

import (
	"github.com/pions/ice"
)

// Configuration defines a set of parameters to configure how the
// peer-to-peer communication via PeerConnection is established or
// re-established.
type Configuration struct {
	// ICEServers defines a slice describing servers available to be used by
	// ICE, such as STUN and TURN servers.
	ICEServers []ICEServer

	// ICETransportPolicy indicates which candidates the ICEAgent is allowed
	// to use.
	ICETransportPolicy ICETransportPolicy

	// BundlePolicy indicates which media-bundling policy to use when gathering
	// ICE candidates.
	BundlePolicy BundlePolicy

	// RTCPMuxPolicy indicates which rtcp-mux policy to use when gathering ICE
	// candidates.
	RTCPMuxPolicy RTCPMuxPolicy

	// PeerIdentity sets the target peer identity for the PeerConnection.
	// The PeerConnection will not establish a connection to a remote peer
	// unless it can be successfully authenticated with the provided name.
	PeerIdentity string

	// Certificates are not supported in the JavaScript/Wasm bindings.
	// Certificates []Certificate

	// ICECandidatePoolSize describes the size of the prefetched ICE pool.
	ICECandidatePoolSize uint8
}

func (c Configuration) getICEServers() (*[]*ice.URL, error) {
	var iceServers []*ice.URL
	for _, server := range c.ICEServers {
		for _, rawURL := range server.URLs {
			url, err := ice.ParseURL(rawURL)
			if err != nil {
				return nil, err
			}
			iceServers = append(iceServers, url)
		}
	}
	return &iceServers, nil
}
