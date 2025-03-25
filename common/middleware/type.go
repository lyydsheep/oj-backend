package middleware

func GetHandlerFunc() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		StartTrace(),
		LogAccess(),
		PanicRecorder(),
	}
}
