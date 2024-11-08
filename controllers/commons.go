package controllers

import (
	"teomebot/clients/streamelements"
	"teomebot/utils"
)

var conDB, _ = utils.OpenDBConnection()

var clientStreamElements = streamelements.NewStreamElementsClient()
