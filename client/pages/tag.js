import React, { useEffect, useState } from "react"
import tagcss from "../styles/tag.module.css"

const tag = () => {
    // const URL_DEV = 'http://www.json-generator.com/api/json/get/coCWKHGGyG?indent=2'
    const URL = 'http://localhost:3030/clickTag/getDetail'

    const GetTagDetail = () => {
        const [Cdata, setCdata] = useState([])
        const [loading, setLoading] = useState(false)
        
        useEffect(() => {
            try {
                async function fetchdata() {
                    const rawdata = await fetch(URL)
                    const jsondata = await rawdata.json()
                    setCdata(jsondata)
                    setLoading(true)
                }
                fetchdata()
            } catch (error) {
                
            }
        },[])
        
        if (loading === true) {
            if (Cdata !== "") {
                return (
                    <div className={tagcss.tag_detailmain}>
                        {Cdata.map((data,idx) => (
                            <div key={idx} className={tagcss.tag_detailcont}>
                                <h3>
                                    <a href={"/clickDetail/"+data.ContentID}>{data.ContentTitle}</a>
                                </h3>
                                <p>{data.ContentContent}</p>
                                <p className={tagcss.tag_tagname} style={{background: data.Tag.TagColor}}>{data.Tag.TagName}</p>
                            </div>        
                        ))}
                    </div>
                )
            } else {
                return (
                    <div>
                        No content sorry :)
                    </div>        
                )
            }
        } else {
            return (
                <div>
                    <div>
                        <div><img src="/loading.gif" draggable="false"/></div>
                    </div>
                </div>
            )
        }
    }

    return (
        <div className={tagcss.tag_main}>
            <a className={tagcss.tag_back} href="/">Back</a>
            {GetTagDetail()}
        </div>
    )
}

export default tag