package main

import (
	ep "github.com/ogame-ninja/extension-patcher"
)

func main() {
	const (
		extensionName                = "tracker"
		webstoreURL                  = "https://chromewebstore.google.com/detail/ogame-tracker/gcebldjabjlagnnnjfodjgiddnonehnd"
		tracker_2024_1_7_1011_sha256 = "3407893e152b94b130f1a22433cfbc7fbe7dbeab864461d1926445f903549994"
	)

	files := []ep.FileAndProcessors{
		ep.NewFile("/manifest.json", processManifest),
	}

	p := ep.MustNew(ep.Params{
		ExtensionName:  extensionName,
		ExpectedSha256: tracker_2024_1_7_1011_sha256,
		WebstoreURL:    webstoreURL,
		Files:          files,
	})
	p.Start()
}

var replN = ep.MustReplaceN

func processManifest(by []byte) []byte {
	by = replN(by, `OGame Tracker`, "OGame Tracker Ninja", 1)
	by = replN(by, `"https://*.ogame.gameforge.com/*"`, `"<all_urls>"`, 2)
	by = replN(by, `"https://*.ogame.gameforge.com/game/*"`, `{old}, "*://*/bots/*/browser/html/*"`, 8)
	return by
}
