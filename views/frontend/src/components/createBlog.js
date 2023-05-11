import { useState } from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import './createBlog.css'
import axios from 'axios'
import Header from './Header'

const CreateBlog = () => {
  const [title, setTitle] = useState('')
  const [author, setAuthor] = useState('')
  const [content, setContent] = useState('')
  const navigate = useNavigate()
  const state = useLocation()

  const handleUserLogin = async (e) => {
    e.preventDefault()
    axios.post("http://localhost:8000/api/blog", {
      title: title,
      author: author,
      content: content,
    }, {
      headers: {
        'Authorization': "Bearer " + localStorage.getItem("token")
      }
    }).then(response => {
      if (response.status === 201) {
        navigate('/home',{ state: {first_name : state.state.first_name } })
      }
    }).catch(err => {
      console.log(err.response.data)
    })
  }

  return (
    <>
      <Header />
      <div className="createDiv">
        <Form>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>Title</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter title"
              value={title}
              onChange={e => {
                setTitle(e.target.value)
              }}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Label>Author</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter name of Author"
              value={author}
              onChange={e => {
                setAuthor(e.target.value)
              }}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="exampleForm.ControlTextarea1">
            <Form.Label>Content</Form.Label>
            <Form.Control
              as="textarea"
              rows={3}
              value={content}
              onChange={e => {
                setContent(e.target.value)
              }}
            />
          </Form.Group>
          <Button
            variant="primary"
            type="submit"
            onClick={e => handleUserLogin(e)}
          >
            Submit
          </Button>
        </Form>
      </div>
    </>
  )
}

export default CreateBlog
