package middleware

//// check if request has client access token
//// on header "ClientAccToken"
//func CheckClientToken(auth IAuth) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// take application context
//		context := ginutil.GetAppCtxOrAbort(c)
//		if context == nil {
//			return
//		}
//		headerToken := c.GetHeader(ginutil.HeaderClientAccToken)
//		if headerToken == "" {
//			log.Info("CheckClientToken", "url no client access token", c.Request.URL)
//			ginutil.AbortWithErr(c, errs.NewUnauthorized().WithMsg("no client token"))
//			return
//		}
//		clientId, err := auth.CheckClientToken(headerToken)
//		if err != nil {
//			ginutil.AbortWithErr(c, err)
//			return
//		} else {
//			context.SetClientId(clientId)
//			c.Next()
//			return
//		}
//	}
//}
