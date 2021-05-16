package httphandler

/*
const (
	// ClientInfoSession session client info.
	ClientInfoSession = "clientInfo"
)
*/
// KeyRequests llaves usadas en las peticiones para el uso de contextos.
const (
	// KeyRequestError error dentro del contexto de gin (gin.Context.Get gin.Context.Set).
	KeyRequestError = "_req_error_"
	// KeyRequestSecret define si la informacion de la peticion debe ser secreta para no realizar el log de la infomacion.
	KeyRequestSecret = "_req_secret_"
	// KeyRequestLogin define si la peticion es debida a un login.
	KeyRequestLogin = "_req_login_"
	// KeySession user session information.
	KeySession = "_req_session_"
	// KeyRequestTx guarda la transaccion de BD (si se ocupó).
	KeyRequestTx = "_req_tx_"
)
