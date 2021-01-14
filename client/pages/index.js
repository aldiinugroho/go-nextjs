import React from "react"
import Nav from "../component/Nav"
import Content from "../component/Content"
import idxcss from "../styles/index.module.css"

const index = () => {
    return (
        <div className={idxcss.content_center_main}>
            <div className={idxcss.content_main}>
                <Nav/>
                <Content/>
            </div>
        </div>  
    )
}

export default index