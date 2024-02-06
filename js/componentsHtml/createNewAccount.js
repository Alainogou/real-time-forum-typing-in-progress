export const createNewAccount=(container) =>{
   
    container.innerHTML = `
    
   
    <form>
        <div><span class="start" >*</span><span class="errorStyle messageErrorNickname"></span></div> 
        <input type="text" id="nickname" name="nickname" placeholder="Nickname">

        <div><span class="start" >*</span><span class="errorStyle messageErrorAge"></span></div> 
        <input type="text" id="age" name="age" placeholder="Age">

        <div><span class="start" >*</span><span class="errorStyle messageErrorGender"></span></div>
        <select id="gender" name="gender">
            <option value="male">Male</option>
            <option value="female">Female</option>
        </select>

        <div><span class="start" >*</span><span class="errorStyle messageErrorFName"></span></div> 
        <input type="text" id="first-name" name="first-name" placeholder="First Name" >

        <div><span class="start" >*</span><span class="errorStyle messageErrorLName"></span></div> 
        <input type="text" id="last-name" name="last-name" placeholder="Last Name" >

        <div><span class="start" >*</span><span class="errorStyle messageErrorEmail"></span></div> 
        <input type="email" id="email" name="email" placeholder="E-mail" >

        <div><span class="start">*</span><span class="errorStyle messageErrorPassword"></span></div> 
        <input type="password" id="password" name="password" placeholder="Password" >
        <input type="password" id="ConfirmPassword" name="ConfirmPassword" placeholder="Confirm password" >
        

        <div class="link submitRegister" >
            <button type="submit" class="login">Register</button>
        </div>   
    </form>
`
    
}
