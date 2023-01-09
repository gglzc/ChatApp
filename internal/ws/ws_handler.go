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


func (h *Handler) CreateRoom (c *gin.Context){
	var  req  CreateRoomReq
	if err:=c.BindJSON(&req) ; err!=nil{
		c.JSON(http.StatusBadRequest , gin.H{"error :":err.Error()})
		return
	}

	h.hub.Rooms[req.ID]=&Room{
		ID: req.ID,
		Name: req.Name,
		Clients: map[string]*Client{},
	}
	c.JSON(http.StatusOK,req)
}
	var upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool{
			return true
		},
	}

func (h *Handler) JoinRoom (c *gin.Context){
	conn,err:=upgrader.Upgrade(c.Writer , c.Request ,nil)
	if err!=nil {
		c.JSON(http.StatusBadRequest , gin.H{"error":err})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userid")
	username := c.Query("username")

	cl:=&Client{
		Conn:     conn,
		Message:  make(chan *Message , 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m:=&Message{
		Content: "A new User Come the room",
		RoomID: roomID,
		Username: username,
	}
	//Register a new client through the register channel
	h.hub.Register<-cl
	//Broadcast the message
	h.hub.Broadcast<-m
	//writeMessage()
	go cl.writeMessage()
	//readMessage()
	cl.readMessage(h.hub)
}

type RoomRes  struct{
	ID			string	`json:"id"`
	Name	    string	`json:"name"`
}

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
func (h *Handler) GetRoomsClients(c *gin.Context){
	var clients []ClientRes
	roomId:=c.Param("roomId")
	//check if there is no room and response non array
	if  _,ok:= h.hub.Rooms[roomId] ; !ok{
		clients=make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}
	//the room exist
	for _,r:=range h.hub.Rooms[roomId].Clients{
		clients = append(clients,ClientRes{
			ID: 	  r.ID,
			UserName: r.Username,
		})
	}
	c.JSON(http.StatusOK , clients)
}