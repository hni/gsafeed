package gsafeed

import (
	"encoding/xml"
	"log"
)

//<?xml version="1.0" encoding="utf-8"?>
//<!DOCTYPE gsafeed PUBLIC "-//Google//DTD GSA Feeds//EN" "">
//<gsafeed>
  //<header>
    //<datasource>web</datasource>
    //<feedtype>incremental</feedtype>
  //</header>
  //<group>
    //<record url="http://www.corp.enterprise.com/hello02" mimetype="text/plain"></record>
  //</group>
//</gsafeed>

type Record struct {
	XMLName    xml.Name `xml:"record"`
	Url        string   `xml:"url,attr"`
	MimeType    string   `xml:"mimetype,attr"`
	CrawlImmediately string   `xml:"crawl-immediately,attr"`
}
type Feed struct {
	XMLName xml.Name `xml:"gsafeed"`
    DataSource string `xml:"header>datasource"`
    FeedType string `xml:"header>feedtype"`
	Records    []Record    `xml:"group>url"`
}

const (
    DocType = `<!DOCTYPE gsafeed PUBLIC "-//Google//DTD GSA Feeds//EN" "">` + "\n"
)

func GsaUrlFeed(urls []string, dataSource string)(string) {
	f := &Feed{}
    f.DataSource = dataSource
    f.FeedType = "metadata-and-url"
    for _, url := range urls {
		f.Records = append(f.Records, Record{Url: url, MimeType: "text/html", CrawlImmediately: "true"})
	}
	output, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

    return xml.Header + DocType + string(output)
}
