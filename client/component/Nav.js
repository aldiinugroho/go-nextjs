import React from "react"
import navcss from "../styles/component_styles/nav_style.module.css"
import limit from "../styles/component_styles/nav_n_cont_style.module.css"
import testdata_tag from "../testdata/testdata_tag"

const Nav = () => {
    const GetTag = () => {
        const URL = testdata_tag
        return (
            <div className={navcss.nav_tagConf}>
                {URL.map((tag,idx) => (
                    <a href="#" key={idx} className={navcss.nav_tag} style={{background: tag.color}}>{tag.tag_name}</a>
                ))}
            </div>
        )
    }
    
    return (
        <div className={navcss.nav_main}>
            <div className={limit.nav_n_cont_limit}>
                {GetTag()}
                {/* <div className={navcss.nav_tagConf}>
                    <div className={navcss.nav_tag}>JavaScript</div>
                    <div className={navcss.nav_tag}>Go</div>
                    <div className={navcss.nav_tag}>Java</div>
                    <div className={navcss.nav_tag}>Ruby</div>
                    <div className={navcss.nav_tag}>Kafka</div>
                    <div className={navcss.nav_tag}>noo gamee</div>
                    <div className={navcss.nav_tag}>lol ting</div>
                    <div className={navcss.nav_tag}>aldi</div>
                </div> */}
                <div className={navcss.nav_mine}>
                    <a href="">Comment</a>
                </div>
                <div className={navcss.nav_dataprof}>
                    <div>Jakarta, Indonesia</div>
                </div>
            </div>
        </div>
    )
}

export default Nav