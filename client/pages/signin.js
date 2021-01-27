import React from "react"
import signincss from "../styles/sign.module.css"

const signin = () => {
    return (
        <div>
            <form className={signincss.signup_form} method="POST" action="/signindata">
                <div>
                    <h1>Lets get inside</h1>
                </div>
                <div><input type="text" placeholder="username" name="username"></input></div>
                <div><input type="password" placeholder="password" name="password"></input></div>
                <div><button>Sign in</button></div>
            </form>
        </div>
    )
}

export default signin