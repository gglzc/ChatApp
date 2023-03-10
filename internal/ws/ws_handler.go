package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct{
	hub *Hub
}

type CreateRoomReq struct{
	ID			string	`json:"id"`
	Name	    string	`json:"name"`
}


func NewHandler (h *Hub) *Handler{
	return &Handler{
		hub: h,
	}
}

// @Summary 創建房間
// @Tags Room API
// @Router /ws/createroom [post]
// @Param id query string true "房間ID"
// @Param name query string true "房間名稱"
// @Success 200 {string} string "創建成功"
func (h *Handler) CreateRoom (c *gin.Context){
	var  req  CreateRoomReq
	if err:=c.ShouldBind(&req) ; err!=nil{
		c.JSON(http.StatusBadRequest , gin.H{"error :":err.Error()})
		return
	}

	h.hub.Rooms[req.ID]=&Room{
		ID: req.ID,
		Name: req.Name,
		Clients: make(map[string]*Client),
	}
	c.JSON(http.StatusOK,req)
}
	
var upgrader = websocket.Upgrader{
	HandshakeTimeout: 0,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	WriteBufferPool:  nil,
	Subprotocols:     []string{},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

// @Summary 加入房間
// @Tags Room API
// @Router /ws/joinRoom/{roomId} [get]
// @Param id query string true "用户 ID"
// @Param roomId path string true "房間 ID"
// @Param username query string true "用户名"
// @Success 200 {string} string "成功"
func (h *Handler) JoinRoom (c *gin.Context){
	conn,err:=upgrader.Upgrade(c.Writer , c.Request ,nil)
	if err!=nil {
		c.JSON(http.StatusBadRequest , gin.H{"error":err})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	client:=&Client{
		Conn:     conn,
		Message:  make(chan *Message),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	message:=&Message{
		Content: "A new User Come the room",
		RoomID: client.RoomID,
		Username: client.Username,
	}
	//Register a new client through the register channel
	h.hub.Register<-client
	//Broadcast the message
	h.hub.Broadcast<-message
	//writeMessage()
	go client.writeMessage()
	//readMessage()
	go client.readMessage(h.hub)
}

type RoomRes  struct{
	ID			string	`json:"id"`
	Name	    string	`json:"name"`
}

// @Summary 列出所有房間
// @Tags Room API
// @Router /ws/getRooms [get]
// @Success 200 {object} []RoomRes "成功"
func (h *Handler) GetRooms (c *gin.Context){
	rooms:=make([]RoomRes , 0)

	for _,r := range h.hub.Rooms{
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	c.JSON(http.StatusOK, rooms)
}
// get client information in the rooms
type ClientRes struct{
	ID				string	`json:"id"`
	UserName	    string	`json:"username"`
}
// @Summary 列出房間中所有user
// @Tags Room API
// @Router /ws/getClients/{roomId} [get]
// @Param roomId path string true "房間 ID"
// @Success 200 {object} []ClientRes "成功"
func (h *Handler) GetRoomsClients(c *gin.Context){
	var clients []ClientRes
	roomId:=c.Param("roomId")
	//check if there is no room and response an empty array
	if  _,ok:= h.hub.Rooms[roomId] ; !ok{
		clients=make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}
	//the room exist
	for _,c:=range h.hub.Rooms[roomId].Clients{
		clients = append(clients,ClientRes{
			ID: 	  c.ID,
			UserName: c.Username,
		})
	}
	c.JSON(http.StatusOK , clients)
}