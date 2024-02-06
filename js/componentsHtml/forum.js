export const headerPage=(container)=>{

    container.innerHTML= `
    <nav>
        <div class="left">
            <div class="logo">
                <h1 class='mnele'>REAL TIME FORUM</h1>
                <!-- <img src="/assets/image/logo.png"> -->
            </div>
        <!-- <div class="search_bar">
            <i class="fa-solid fa-magnifying-glass"></i>
            <input type="text" placeholder="Search">
        </div> -->
                    
        </div> 
        <div class="right">
        
            <div class="nav-connection">
                <!-- <i class="fa-solid fa-list-ul"></i> -->
                <i class="fa-brands fa-facebook-messenger" id = "messageHeader"></i>
                <i class="fa-solid fa-user" id = "userHeader"></i>
                <i  class="fa-solid fa-right-from-bracket" id ="logoutHeader"></i>
                <!-- <img src="/assets/images/profil.png"> -->
            </div>         
        </div>
    </nav>
    
    `
}


export const displayCategories=(container, firstName, lastName)=>{
    container.innerHTML=`
        <div class="left">
        <div>
            <div class="img">
                <div class="user" style="background-color:#efefef; height:30px;width:30px; text-align:center; border-radius:50%; padding-top:4px">
                 <i class="fa-solid fa-user" ></i>
                 </div>
                <p>${firstName} ${lastName}</p>
            </div>
            <hr>
        </div>
    
        <h2>Categories</h2>
        
        <div class="shortcuts">
            <img src="/assets/image/sport.avif">
            <p>Sport</p>
        </div>
        <div class="shortcuts">
            <img src="/assets/image/art.avif">
            <p>Art</p>
        </div>
        <div class="shortcuts">
            <img src="/assets/image/informatique.avif">
            <p>Informatics</p>
        </div>
        <div class="shortcuts">
            <img src="/assets/image/religion.avif">
            <p>Religion</p>
        </div>
        <div class="shortcuts">
            <img src="/assets/image/game.avif">
            <p>Game</p>
        </div>
        
    </div>`
}

export const Messenger = (container, nickname , status, NombreMessage) => {
   let div = document.createElement('div');
   div.className = 'third_warpper';
   div.style.marginBottom = '10px'
   let contactDiv = document.createElement('div');
   contactDiv.classList.add( `contact-${nickname}`);
   contactDiv.classList.add( 'contact');
   contactDiv.setAttribute('data-userid', `${nickname}`)
  
//    contactDiv.style.justifyContent='space-between'
 
   
   let userDiv = document.createElement('div');
   userDiv.className = 'user';
   userDiv.style.backgroundColor = '#efefef';
   userDiv.style.height = '30px';
   userDiv.style.width = '30px';
   userDiv.style.textAlign = 'center';
   userDiv.style.borderRadius = '50%';
   userDiv.style.paddingTop = '4px';
   let i = document.createElement('i');
   i.className = 'fa-solid fa-user';
   userDiv.appendChild(i);

   let iconeMesage = document.createElement('div');
   iconeMesage.className = 'iconeMessage';
   iconeMesage.style.backgroundColor = '#efefef';
   iconeMesage.style.height = '30px';
   iconeMesage.style.width = '30px';
   iconeMesage.style.textAlign = 'center';
   iconeMesage.style.borderRadius = '50%';
   iconeMesage.style.paddingTop = '4px';
   let classI= document.createElement('i');
   classI.className = "fa-solid fa-message";
   iconeMesage.appendChild(classI);

   let countMess = document.createElement('span');
   countMess.id=(`Nmessage-${nickname}`);
   countMess.textContent = NombreMessage

   countMess.style.color="red";
   let leftContactDiv = document.createElement('div');
   leftContactDiv.className ="leftContact";

   
   leftContactDiv.appendChild(iconeMesage)
   leftContactDiv.appendChild(countMess);
   leftContactDiv.appendChild(userDiv);

   
   let p = document.createElement('p');
    p.innerText = nickname + " ";
    leftContactDiv.appendChild(p)
    let rightContactDiv = document.createElement('div');
    rightContactDiv.className ="rightContact-"+nickname;
    let span = document.createElement('span');
    span.id = "status-" + nickname; // Remplacer 'nickname' par une valeur unique si nécessaire
    span.innerText = status;
    span.style.display="none";
    if (status === 'offLine') {
        span.classList.add('status-offline');
           // Ajoutez une image si le statut est 'offline'
            let img = document.createElement('img');
            img.src = 'assets/image/status-no-active-svgrepo-com.svg'; // Remplacez par le chemin de votre icône
            img.alt = 'Offline';
            img.className = 'status-icon-off'; // Ajoutez une classe pour styliser l'image si nécessaire
            rightContactDiv.appendChild(img);
       }else{
        span.classList.add('status-online');
        let img = document.createElement('img');
        img.src = 'assets/image/status-active-svgrepo-com.svg'; // Remplacez par le chemin de votre icône
        img.alt = 'Online';
        img.className = 'status-icon-on'; // Ajoutez une classe pour styliser l'image si nécessaire
        rightContactDiv.appendChild(img);
       }

       contactDiv.appendChild(leftContactDiv);

   
    contactDiv.appendChild(span);
    contactDiv.appendChild(rightContactDiv);
   
    container.appendChild(contactDiv);

   
};



