import React, { useState } from 'react'
import './Home.css'
import Card from 'react-bootstrap/Card'
import axios from 'axios'
import { RiDeleteBinFill } from 'react-icons/ri'
import Button from 'react-bootstrap/Button'
import Header from './Header'

const Home = () => {
  const [allBlogData, setAllBlogData] = useState([])
  const [loading, setLoading] = useState(false);

  React.useEffect(() => {
    getAllBlogs()
  }, [])

  const getAllBlogs = () => {
    setLoading(true);
    axios.get('http://localhost:8000/api/blog',{
      headers: {
        'Authorization' : "Bearer " + localStorage.getItem("token")
      }
    }).then(response => {
      const allBlog = response.data.data
      if (response.status === 200) {
        setAllBlogData(allBlog);
        setLoading(false);
      }
    })
  }

  const deleteBlog = (id) => {
    axios.delete(`http://localhost:8000/api/blog/${id}`,{
      headers: {
        'Authorization' : "Bearer " + localStorage.getItem("token")
      }
    }).then(response => {
      if (response.status === 200)  {
        getAllBlogs()
      }
    })
  }

  return (
    <>
      <Header />
      {loading ? <h3 style={{ color: 'white' }}>Loading..</h3> :
        <>
          {allBlogData.length > 0 ?
            <div className="homeDiv">
              {allBlogData.map((blog) => {
                return (
                  <Card className="HomeCard" key={blog.id}>
                    <Card.Header>
                      Author - {blog.author}{' '}
                      <Button
                        variant="outline-dark"
                        onClick={() => deleteBlog(blog.id)}
                        style={{ float: 'right' }}
                      >
                        <RiDeleteBinFill />
                      </Button>{' '}
                    </Card.Header>
                    <Card.Body>
                      <Card.Title>Title :- {blog.title}</Card.Title>
                      <Card.Text>{blog.content}</Card.Text>
                    </Card.Body>
                  </Card>
                )
              })}
            </div>
            : <h3 style={{ textAlign: 'center', color: 'white' }}>No Blogs Found</h3>
          }
        </>
      }
    </>
  )
}

export default Home
