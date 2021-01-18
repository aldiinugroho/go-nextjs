import React from "react"
import inpcss from "../styles/input.module.css"

const input = () => {
    return (
        <div className={inpcss.detail_main}>
            <div className={inpcss.input_title}>Please consider language while inputing data !!</div>
            <div>
                <form className={inpcss.detail_mainform} action="/inputdata" method="POST">
                    <div><input type="text" placeholder="Title" name="title"></input></div>
                    <div><textarea name="message" rows="5" cols="30" placeholder="Content"></textarea></div>
                    <div><input type="text" placeholder="Tag name" name="tag"></input></div>
                    <div><input type="text" placeholder="Tag color" name="tagcolor"></input></div>
                    <div><button>Submit</button></div>
                </form>
            </div>
        </div>
    )
}

export default input