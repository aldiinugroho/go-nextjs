import React from "react"
import errcss from "../styles/errorcss.module.css"

const errorpage =  () => {
    
    return (
        <div className={errcss.error_main}>
            <div className={errcss.error_handle}>
                <div><img src="/errimg1.png" draggable="false" alt="errorimg"/></div>
                <div><span className={errcss.error_span404}>404</span> not found.</div>
            </div>
            <div className={errcss.error_btm}>
                {/* <div>©</div> */}
                <div className={errcss.error_comt}>
                    <a href="">Comment</a>
                </div>
            </div>
        </div>
    )
}

export default errorpage