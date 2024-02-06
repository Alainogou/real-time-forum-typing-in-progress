
let app=document.getElementById('connexion')


let ap=document.getElementById('enter')



import {createNewAccount} from './componentsHtml/createNewAccount.js'
import {Messenger,  displayCategories, headerPage} from './componentsHtml/forum.js'
import { sendForm } from './componentsHtml/loginForm.js'
import {createPostbutton, postForm} from './componentsHtml/postForm.js'

import { chatContainer } from './hanleJs/chat.js'

import { fetchPost, handleCreatePost } from './hanleJs/post.js'
import { handleRegistration } from './hanleJs/register.js'
import { alertMessage } from './helper/alert.js'
import { moveUserToTop } from './helper/movetoUp.js'



document.addEventListener('DOMContentLoaded', (event) => {
    
    fetch('http://localhost:8081/auth')
    .then(response => response.json())
    .then(data => {
        

        if (data.IsAuth){
            handleSuccessfulLogin(data)

        }else{
            Authentifaction(ap)
           
         
        }   
       
    
    })
    .catch(error => console.error('Erreur:', error));




});

function Authentifaction(ap){
    sendForm(ap)
    let loginForm=document.getElementById('loginForm')

    let registrationForm=document.getElementById('registrationForm')
    const homeView = document.querySelector(".home-view");
    
    let register= document.querySelector(".registration")
    let closeForm=document.querySelector("#close-register-form")
 
    if (registrationForm) registrationForm.addEventListener('submit', handleRegistration);
    if (loginForm) loginForm.addEventListener('submit', handleLogin);
     
     let creatNewacc=document.querySelector(".button-new-account")

     if (creatNewacc) creatNewacc.addEventListener("click", function(){
         homeView.style.display="none"
         register.style.display="block"
         createNewAccount(registrationForm)
     })
     

     if (closeForm) closeForm.addEventListener("click", function(){
         register.style.display="none"
         homeView.style.display="flex"
       
        
     })
}







function handleSuccessfulLogin(data) {
   
    ap.style.display="none";
    headerPage(app);
    let main=document.createElement('div');
    let right = document.createElement('div');
    let center=document.createElement('div')
    center.classList.add('center');

    right.classList.add('right');
    right.classList.add('right1');

    main.classList.add('main');
    

    let globalPosts= document.createElement('div')
    globalPosts.classList.add('mainPost')
   
    displayCategories(main, data.User.FirstName, data.User.LastName);
    createPostbutton(center, data.User.NickName)
   
   
    fetchPost(globalPosts,data.User.Id)
   
    let postform= document.createElement('div')

    center.appendChild(globalPosts)
    main.appendChild(center)
    app.appendChild(postform)
    app.appendChild(main);
    

    let showPostForm= document.querySelector(".showPostForm")
    if (showPostForm) showPostForm.addEventListener("click", function(event){
        postform.style.display='block'

        postform.style.position = "relative"
        postform.style.top = "0px"
        postForm(postform, data.User.Id)

        let closeForm=document.querySelector(`.btn-close`)
       
        
        if (closeForm) closeForm.addEventListener("click", function(){
            postform.style.display='none'
          
       })

       let postForms=document.querySelector("#postForm")
      
       if (postForms) {
            postForms.addEventListener('submit', function(event) {
                handleCreatePost(event, postform);
            });
       }
      
    })


    let logoutHeader=document.getElementById("logoutHeader");
    if (logoutHeader) logoutHeader.addEventListener("click",()=>{
        logout(ap, data.User.NickName);
    });

    
    const socket = new WebSocket('ws://localhost:8081/ws');

    socket.onopen = (event) => {
        let message = JSON.stringify({NickName: data.User.NickName});
        socket.send(message);
    };

   
    socket.onmessage = function(event) {
        let msg = JSON.parse(event.data);
        console.log(msg);
        

        if (msg.NewConnection){
            let userConnected=document.querySelector(`.rightContact-${msg.PersonConnected}`)
            if (userConnected) {
            userConnected.innerHTML=''
            let img = document.createElement('img');
            img.src = 'assets/image/status-active-svgrepo-com.svg'; 
            img.alt = 'Online';
            img.className = 'status-icon-on'
            userConnected.appendChild(img);
                
            }
        }else if (msg.NewDeconnexion){
            let userConnected=document.querySelector(`.rightContact-${msg.PersonConnected}`)
            if (userConnected) {
            userConnected.innerHTML=''
            let img = document.createElement('img');
            img.src = 'assets/image/status-no-active-svgrepo-com.svg'; // Remplacez par le chemin de votre icône
            img.alt = 'Offline';
            img.className = 'status-icon-off'
             userConnected.appendChild(img)
            }//  PersonConnected string
            
        }else if  (msg.NewMessage===true){    
           
            let nMsgElement = document.querySelector(`#Nmessage-${msg.PersonConnected}`);
           
            if (nMsgElement) {
                nMsgElement.innerText= parseInt(nMsgElement.textContent) + 1 
            }
             moveUserToTop(msg.PersonConnected)
            setTimeout(()=>{
               let body= document.querySelector('body')
               alertMessage(msg.PersonConnected,data.User.NickName,body)  
            },1000)
        }else{

            right.innerHTML=''
           
            let div = document.createElement('div');
            div.className = 'third_warpper';
            
            let contactTagDiv = document.createElement('div');
            contactTagDiv.classsName = 'contact_tag';
            let h2 = document.createElement('h2');
            h2.innerText = 'Contacts';
            contactTagDiv.appendChild(h2);
            div.appendChild(contactTagDiv);

            let userlist = document.createElement("div");
            userlist.className = "userlist";
            let AllUser=[...msg.MessageExist, ...msg.NotMessage]
            for (let k=0;k<AllUser.length;k++){
                if (AllUser[k].NickName !== data.User.NickName){
                    Messenger(userlist, AllUser[k].NickName, AllUser[k].Status, AllUser[k].NbreMessages )
                    
                    setTimeout(() => {
                        let contact = document.querySelector(`.contact-${AllUser[k].NickName}`)
                        
                        if (contact) contact.addEventListener("click",()=>{
                            console.log("contact clicked");
                            let premierDive= document.querySelector(`.chat-card-${AllUser[k].NickName}`)
                            if (premierDive) premierDive.remove()

                            let premierDiv = document.createElement('div');
                            premierDiv.classList.add('chat-card');
                            premierDiv.classList.add(`chat-card-${AllUser[k].NickName}`);

                            let messageOpen=document.querySelector(`#Nmessage-${AllUser[k].NickName}`)
                            if (messageOpen) messageOpen.innerHTML="0"
                            
                            
                            chatContainer(AllUser[k].NickName, data.User.NickName, premierDiv)
                            right.appendChild(premierDiv)
                        
                        })

                    
                    }, 1000);

                }
               
            }
            div.appendChild(userlist)
            right.appendChild(div)
        }
        
        
 
       
        
      
    };
           

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    socket.onerror = (error) => {
        console.log(`WebSocket error: ${error}`);
    };

   
    main.appendChild(right);
}






async function handleLogin(event) {
    event.preventDefault();
    const formData = new FormData(event.target);

    let logRequest = {
        EmailOrUsername: formData.get("email-nickname"),
        Password: formData.get("password")
    }

    try {
        const response = await fetch('http://localhost:8081/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(logRequest),
        });

        if (response.ok) {
         
            ap.style.display="none"
            app.style.display="block"
            fetch('http://localhost:8081/auth')
            .then(response => response.json())
            .then(data => {
                if (data.IsAuth){
                    handleSuccessfulLogin(data) 
                }  
            
            })
            .catch(error => console.error('Erreur:', error));
            
        } else {
            const data = await response.json();
            let logNotMatch = document.querySelector(".logNotMatch");
            
            if (data['error_class'] === "logNotMatch") {
                logNotMatch.innerHTML = data['message']
            }
            

            setTimeout(function () {
                logNotMatch.innerHTML = ''
            }, 5000);
        }
    } catch (error) {
        console.error('Erreur lors de la création de l\'utilisateur:', error);
    }
}


async function logout(ap, userName) {
    let userDeconn = {
      NickName: userName,
    }
    try {
        const response = await fetch('http://localhost:8081/logout', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userDeconn),
        });

        if (response.ok) {
            
            app.style.display="none"
            ap.style.display="block"
            sendForm(ap)
            let loginForm=document.getElementById('loginForm')

            let registrationForm=document.getElementById('registrationForm')
            const homeView = document.querySelector(".home-view");
            
            let register= document.querySelector(".registration")
            let closeForm=document.querySelector("#close-register-form")
            
            registrationForm.addEventListener('submit', handleRegistration);
            loginForm.addEventListener('submit', handleLogin);
            
            document.querySelector(".button-new-account").addEventListener("click", function(event){
                homeView.style.display="none"
                register.style.display="block"
                createNewAccount(registrationForm)
            })

            closeForm.addEventListener("click", function(){
                register.style.display="none"
    
                homeView.style.display="flex"
              
               
            })


            // loadConnexionPage(app)
        } else {
           
        
        }
    } catch (error) {
        console.error('Erreur lors de la création de l\'utilisateur:', error);
    }
}