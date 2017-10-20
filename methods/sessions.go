/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - nethserver@nethesis.it
 *
 * This script is part of NethServer.
 *
 * NethServer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * NethServer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with NethServer.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

 package methods

 import (
	 "net/http"
	 "time"

	 "github.com/gin-gonic/gin"
	 _ "github.com/jinzhu/gorm/dialects/mysql"

	 "sos-server/database"
	 "sos-server/models"
 )

 func CreateSession(c *gin.Context) {
	 sessionId := c.PostForm("session_id")
	 lk := c.PostForm("lk")
	 started := time.Now().String()

	 session := models.Session{
		 Lk: lk,
		 SessionId: sessionId,
		 VpnIp: "",
		 Started: started,
	 }

	 db := database.Database()
	 db.Save(&session)

	 c.JSON(http.StatusCreated, gin.H{"id": session.Id})
 }

 func UpdateSession(c *gin.Context) {
	 var session models.Session
	 lk := c.Param("lk")
	 vpnIp := c.PostForm("vpn_ip")

	 db := database.Database()
	 db.Where("lk = ?", lk).First(&session)

	 if session.Id == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		 return
	 }

	 session.VpnIp = vpnIp
	 db.Save(&session)
 }

 func GetSessions(c *gin.Context) {
	 var sessions []models.Session

	 db := database.Database()
	 db.Find(&sessions)

	 if len(sessions) <= 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No sessions found!"})
		 return
	 }

	 c.JSON(http.StatusOK, sessions)
 }

 func GetSession(c *gin.Context) {
	 var session models.Session
	 sessionId := c.Param("session_id")

	 db := database.Database()
	 db.Where("session_id = ?", sessionId).First(&session)

	 if session.Id == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		 return
	 }

	 c.JSON(http.StatusOK, session)
 }

 func DeleteSession(c *gin.Context) {
	 var session models.Session
	 var history models.History
	 lk := c.Param("lk")

	 db := database.Database()
	 db.Where("lk = ?", lk).First(&session)

	 if session.Id == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		 return
	 }

	 // add to history this session
	 history.SessionId = session.SessionId
	 history.Lk = session.Lk
	 history.Started = time.Now().String()
	 db.Save(&history)

	 // delete session
	 db.Delete(&session)

	 c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully!"})
 }