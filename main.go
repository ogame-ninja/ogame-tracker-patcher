package main

import (
	ep "github.com/ogame-ninja/extension-patcher"
)

func main() {
	const (
		webstoreURL                    = "https://chromewebstore.google.com/detail/ogame-tracker/gcebldjabjlagnnnjfodjgiddnonehnd"
		tracker_2024_7_12_17090_sha256 = "d462665a0c6c961c3616c957112b80894b638050e06858ce0cdf65e50b0399c1"
	)

	files := []ep.FileAndProcessors{
		ep.NewFile("manifest.json", processManifest),
		ep.NewFile("service-worker.js", processServiceWorker),
		ep.NewFile("content-scripts/message-tracking.js", processMessageTracking),
		ep.NewFile("views/js/chunk-common.066aa981.js", processChunkCommon),
	}

	ep.MustNew(ep.Params{
		ExpectedSha256: tracker_2024_7_12_17090_sha256,
		WebstoreURL:    webstoreURL,
		Files:          files,
	}).Start()
}

var replN = ep.MustReplaceN

func processManifest(by []byte) []byte {
	by = replN(by, `OGame Tracker`, "OGame Tracker Ninja", 1)
	by = replN(by, `"https://*.ogame.gameforge.com/*"`, `"<all_urls>"`, 2)
	by = replN(by, `"https://*.ogame.gameforge.com/game/*"`, `{old}, "*://*/bots/*/browser/html/*"`, 8)
	return by
}

func processServiceWorker(by []byte) []byte {
	return by
}

func processMessageTracking(by []byte) []byte {
	by = replN(by, `/game/index.php`, ``, 1)
	return by
}

func processChunkCommon(by []byte) []byte {
	by = replN(by, `*://*.ogame.gameforge.com/*`, `*://*/bots/*/browser/html/*`, 1)
	return by
}
