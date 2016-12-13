package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/etcenter/c4/asset"
)

const version_number = "0.6.8"

func main() {
	words := []string{
		"alfa",    //c43zYcLni5LF9rR4Lg4B8h3Jp8SBwjcnyyeh4bc6gTPHndKuKdjUWx1kJPYhZxYt3zV6tQXpDs2shPsPYjgG81wZM1
		"bravo",   //c42jd8KUQG9DKppN1qt5aWS3PAmdPmNutXyVTb8H123FcuU3shPxpUXsVdcouSALZ4PaDvMYzQSMYCWkb6rop9zhDa
		"charlie", //c44erLietE8C1iKmQ3y4ENqA9g82Exdkoxox3KEHops2ux5MTsuMjfbFRvUPsPdi9Pxc3C2MRvLxWT8eFw5XKbRQGw
		"delta",   //c42Sv2Wi2Qo8AKbJKnUP6YTSdz8pt9aDaf2Ltx44HF1UDdXANM8Ltk6qEzpncvmVbw6FZxgBumw9Eo2jtGyaQ5gDSC
		"echo",    //c41bviGCyTM2stoMYVTVKgBkfC6SitoLRFinp77BcmN9awdaeC9cxPy4zyFQBhmTvRzChawbECK1KBRnw3KnagA5be
		"foxtrot", //c427CsZdfUAHyQBS3hxDFrL9NqgKeRuKkuSkxuYTm26XG7AKAWCjViDuMhHaMmQBkvuHnsxojetbQU1DdxHjzyQw8r
		"golf",    //c41yLiwAPdsjiBAAw8AFwQGG3cAWnNbDio21NtHE8yD1Fh5irRE4FsccZvm1WdJ4FNHtR1kt5kev7wERsgYomaQbfs
		"hotel",   //c44nNyaFuVbt5MCfo2PYWHpwMkBpYTbt14C6TuoLCYH5RLvAFLngER3nqHfXC2GuttcoDxGBi3pY1j3pUF2W3rZD8N
		"india",   //c41nJ6CvPN7m7UkUA3oS2yjXYNSZ7WayxEQXWPae6wFkWwW8WChQWTu61bSeuCERu78BDK1LUEny1qHZnye3oU7DtY
	}
	var ids asset.IDSlice

	for i, word := range words {
		e := asset.NewIDEncoder()
		_, err := io.Copy(e, bytes.NewReader([]byte(word)))
		if err != nil {
			panic(err.Error())
		}
		id := e.ID()
		fmt.Printf("%d: \"%s\", %s\n", i, word, id.String())
		ids.Push(id)
	}

	ids_chan := asset.Tree(ids)

	for id := range ids_chan {
		_ = id
		fmt.Printf("\n\nResult: %s\n", id.String())
	}
}
