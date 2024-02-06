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