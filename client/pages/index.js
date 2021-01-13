import React from "react";
import idxcss from "../styles/Home.module.css"
const index = () =>{
  return (
      <div>
        <h1 className={idxcss.index_main}>Hello world, i am Aldi Nugroho</h1>
        <form className={idxcss.index_form} action="/menu" method="POST">
          <div><input type="text" name="username" placeholder="username"/></div>
          <div><input type="number" name="age" placeholder="age"/></div>
          <div><button>Submit</button></div>
        </form>
        <a href="/menu">menu</a>
      </div>
  )
}

export default index