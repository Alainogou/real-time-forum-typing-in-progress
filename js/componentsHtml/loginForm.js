


export function sendForm(ap){
    let form=document.createElement('div')
            form.innerHTML=`<div class="container flex"  id="part1"  >
            <div class="home-view flex">
              <div class="text">
                <h1>REAL TIME FORUM</h1>
                <p>Connect, Collaborate, Communicate<br>  
                   Real-Time Conversations Unleashed! </p>
              </div>
              
              <div class="loginForm" id="login">
                <div><span class="errorStyle logNotMatch"></span></div> 
                <form id="loginForm">
                    <input type="text"  name="email-nickname" placeholder="Nickname or email" >
                    <input type="password"  name="password" placeholder="Password">
                    <div class="link">
                      <button type="submit" class="login">Login</button>
                    </div>
                </form>
             
                <hr>
                <div class="button-new-account">
                    <button type="submit" class="login">Create new account</button>
                </div>
              </div>
            </div>
        
          </div>
        
          <div class="registration" style="display: none;" id="part2">
            <button class="close" id="close-register-form">X</button>
            <form id="registrationForm">
              
            </form>
          </div>
         `
    ap.appendChild(form)

}