import "time"
import "list"
{
    values: [...[string & != "", string & != "", string, time.Format("2006-01-02T15:04:05-0700")]]

    shortnames: [for x in values { x[0] }]
    urls: [for x in values { x[1] }]


  _shortnamesunique: true & list.UniqueItems(shortnames)
  _urlsunique: true & list.UniqueItems(urls)
}


