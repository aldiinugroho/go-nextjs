import React, { useEffect, useState } from "react"
import cont from "../styles/component_styles/content_style.module.css"
import limit from "../styles/component_styles/nav_n_cont_style.module.css"

const Content = () => {
    const URL = 'http://www.json-generator.com/api/json/get/bMRCXUIhhK?indent=2'

    const GetData = () => {
        const [Cdata, setCdata] = useState([])
        
        useEffect(() => {
            try {
                async function fetchdata() {
                    const rawdata = await fetch(URL)
                    const jsondata = await rawdata.json()
                    setCdata(jsondata)
                }
                fetchdata()
            } catch (error) {
                
            }
        },[])
        if (Cdata !== "") {
            return (
                <div>
                    {Cdata.map((data,idx) => (
                        <div key={idx} className={cont.cont_container}>
                            <h3>
                                <a href="/detail">{data.title}</a>
                            </h3>
                            <p>{data.content}</p>
                            <p className={cont.cont_tag}>{data.tag}</p>
                        </div>        
                    ))}
                </div>
            )
        } else {
            return (
                <div className={cont.cont_container}>
                    No content sorry :)
                </div>        
            )
        }
    }

    return (
        <div className={cont.cont_main}>
            <div className={limit.nav_n_cont_limit}>
                <div className={cont.imagetopConf}>
                    <img src="/top.png" draggable="false"/>
                </div>
                {GetData()}
                {/* <div className={cont.cont_container}>
                    <h3>          
                        <a href="#">Parse json in javascript</a>
                    </h3>
                    <p>Lorem ipsum, or lipsum as it is sometimes known, is dummy text used in laying out print, 
                    graphic or web designs. The passage is attributed to an unknown typesetter in the 15th century who is thought to have scrambled parts of Cicero's De Finibus Bonorum et Malorum for use in a type specimen book. Lorem ipsum, 
                    or lipsum as it is sometimes known, is dummy text used in laying out print, graphic or web designs. 
                    The passage is attributed to an unknown typesetter in the 15th century who is thought to have scrambled parts of Cicero's De Finibus Bonorum et Malorum for use in a type specimen book. 
                    Lorem ipsum, or lipsum as it is sometimes known, 
                    is dummy text used in laying out print, graphic or web designs. 
                    The passage is attributed to an unknown typesetter in the 15th century who is thought to have scrambled parts of Cicero's De Finibus Bonorum et Malorum for use in a type specimen book. Lorem ipsum, 
                    or lipsum as it is sometimes known, is dummy text used in laying out print, graphic or web designs. 
                    The passage is attributed to an unknown typesetter in the 15th century who is thought to have scrambled parts of Cicero's De Finibus Bonorum et Malorum for use in a type specimen book.
                    </p>
                    <p className={cont.cont_tag}>JavaScript</p>
                </div>         */}
                
            </div>
        </div>
    )
}

export default Content