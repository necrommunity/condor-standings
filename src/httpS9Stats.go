package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

func httpS9Stats(c* gin.Context) {
	
	data := TemplateData{
		Title: 			 "Season 9 BADS",
	}
	httpServeTemplate(w, "s9stats", data)
}