package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"fmt"
	"os/exec"
	"time"
	"strings"
)

func InitVPN(sessionId string, lk string) string {
	out, err := exec.Command("go","version").Output()
	if err != nil {
	  fmt.Printf("%s", err)
	}
	return strings.TrimSpace(string(out[:]))
}

func CreateSession(c *gin.Context) {
	sessionId := c.PostForm("session_id")
	lk := c.PostForm("lk")
	vpnIp := InitVPN(sessionId, lk)
	started := int(time.Now().Unix())

	session := Session{
		Lk: lk,
		SessionId: sessionId,
		VpnIp: vpnIp,
		Started: started,
	}

	db := Database()
	db.Save(&session)

	c.JSON(http.StatusCreated, gin.H{"id": session.Id})
}

func GetSessions(c *gin.Context) {
	var sessions []Session

	db := Database()
	db.Find(&sessions)

	if len(sessions) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No sessions found!"})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

func GetSession(c *gin.Context) {
	var session Session
	sessionId := c.Param("session_id")

	db := Database()
	db.Where("session_id = ?", sessionId).First(&session)

	if session.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		return
	}

	c.JSON(http.StatusOK, session)
}

func DeleteSession(c *gin.Context) {
	var session Session
	sessionId := c.Param("session_id")

	db := Database()
	db.Where("session_id = ?", sessionId)

	if session.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		return
	}
	db.Delete(&session)

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully!"})
}