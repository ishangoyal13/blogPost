import React, { useState } from 'react'
import './Home.css'
import Card from 'react-bootstrap/Card'
import axios from 'axios'
import {RiDeleteBinFill} from 'react-icons/ri'

const Home = () => {
  const [data, getBlog] = useState('')

  React.useEffect(() => {
    getAllBlogs()
  },[])

  const getAllBlogs = () => {
    axios
      .get("http://localhost:8080/blog")
      .then((response) => {
        const allBlog = response.data.data
        getBlog(allBlog)
      });
  }
  return (
    <DisplayBlogs data={data} />
  )
}

const DisplayBlogs = (props) => {
  const {data} = props
  if (data.length > 0) {
    return (
      <div className='homeDiv'>
        {data.map((note, index) => {
          return <Card className='HomeCard' key={index}>
            <Card.Header>Feature no.- {note.id} <RiDeleteBinFill /></Card.Header>
            <Card.Body>
              <Card.Title>Title :- {note.title}</Card.Title>
              <Card.Text>
                {note.content}
              </Card.Text>
            </Card.Body>
          </Card>
        })}
      </div>
    )
  } else {
    <div>No blog</div>
  }
} 

export default Home