import React from "react"
import signupcss from "../styles/sign.module.css"

const signup = () => {
    return (
        <div>
            <form className={signupcss.signup_form} method="POST" action="/signupdata">
                <div>
                    <h1>Create your account</h1>
                </div>
                <div><input type="text" placeholder="username" name="username"></input></div>
                <div><input type="text" placeholder="email" name="email"></input></div>
                <div><input type="password" placeholder="password" name="password"></input></div>
                <div><button>Create account</button></div>
            </form>
        </div>
    )
}

export default signup