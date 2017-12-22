package fake

var EasypostFakeBatch string

func init() {
	EasypostFakeBatch = `
		{
      "id": "batch_3dc0df9695654f1f8a788e1c872a0769",
      "label_url": null,
      "mode": "test",
      "num_shipments": 1,
      "object": "Batch",
      "reference": null,
      "scan_form": null,
      "shipments": [
        {
          "id": "shp_3dc0df9695654f1f8a788e1c872a0769",
          "batch_status": "postage_purchased",
          "batch_message": null,
          "reference": null
        }
      ],
      "state": "created",
      "status": {
        "created": 0,
        "queued_for_purchase": 0,
        "creation_failed": 0,
        "postage_purchased": 0,
        "postage_purchase_failed": 0
      },
      "label_url": null,
      "created_at": "2014-07-22T07:34:39Z",
      "updated_at": "2014-07-22T07:34:39Z"
    }
	`
}