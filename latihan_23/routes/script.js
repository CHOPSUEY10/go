            

        


        class ChatEvent {

            constructor(type,payload){
                this.type = type
                this.payload = payload
            }


            routeEvent(){
                if (this.type === undefined){
                    alert("No type field in the event")
                }

                // FIX IT LATER FOR NEW FEATURES
            //     switch(this.type){
            //         case "new message" :
            //             console.log("new message")
            //             break
            //         default : 
            //             alert("unsupported message type")
            //             break     
            //     }
            }

            static sendEvent(eventName, payload){
                const event = new ChatEvent(eventName,payload)
                conn.send(JSON.stringify(event))
            }

        }

            const chatMessages = document.querySelector('#chatmessages');
            const MAX_LINES = parseInt(chatMessages.getAttribute('rows')) || 10;

            let selectedChat ="general"
            let conn = null
            const changeChatRoom  = (e) => {
                 e.preventDefault();
                let newchat = document.querySelector("#chatroom")
                if(newchat !== null && newchat.value != selectedChat){
                    console.log(newchat);
                }
                
                return false
                
            }   
            
            const sendMessage = (e) => {
                e.preventDefault();
                let newmessage = document.querySelector("#message")
            
                if(newmessage != null){
                    ChatEvent.sendEvent("send message",newmessage.value)
                }
                return false

            }

            const streamMessage =  (e) => {
            
                // A. Ambil isi teks yang sudah ada (termasuk isinya jika belum penuh)
                let newMessage =  e.payload
                let currentContent = chatMessages.value;


                // B. Tambahkan pesan baru di baris terbawah (APPEND)
                // JIKA textarea sudah ada isinya, tambahkan baris baru (\n) dulu
                if (currentContent !== "") {
                    chatMessages.value = currentContent + '\n' + newMessage;
                } else {
                    // JIKA masih kosong (misalnya baru connect), langsung isi
                    chatMessages.value = newMessage;
                }
                
                // --- Jaga Agar Scroll Di Paling Bawah ---
                chatMessages.scrollTop = chatMessages.scrollHeight;
            }

            const login = async (e) => {
                e.preventDefault();

                let formData = {
                    "username": document.querySelector("#username").value,
                    "password": document.querySelector("#password").value,
                }

                try {
                    const res = await fetch("/login", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json"
                        },
                        body: JSON.stringify(formData),
                        mode: "cors"
                    })
                    if (!res.ok) {
                        const errorText = await res.text()
                        throw errorText || "Unauthorized"
                    }
                    const data = await res.json()

                    connectWebsocket(data.otp)

                } catch (err) {
                    console.error(err)
                }
            }

            const connectWebsocket = (otp) =>{
                
                if(window["WebSocket"]){
                    
                    console.log("support websockets")
                    conn = new WebSocket(`ws://${location.host}/ws?otp=${otp}`);
                    
                    conn.onmessage = (e) => {
                        const eventData = JSON.parse(e.data)
                        const event = Object.assign(new ChatEvent(),eventData)
                        event.routeEvent()    
                        streamMessage(eventData)
                    }

                    conn.onopen = () => {
                        document.querySelector("#connection-header").innerHTML = "Connected to Chat"
                    }

                    conn.onerror = (e) => console.log("ERROR", e);

                    conn.onclose = (e) =>{
                          document.querySelector("#connection-header").innerHTML = "Disconnected"
                    }


                }else{
                    alert("Browser does not support websocket")
                }
            }




            window.onload = () => {
                document.querySelector("#chatroom-selection").onsubmit = changeChatRoom
                document.querySelector("#chatroom-message").onsubmit = sendMessage
                document.querySelector("#login-form").onsubmit = login
            }


