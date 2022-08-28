import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import './createBlog.css'

const CreateBlog = () => {
  const [title, setTitle] = useState('')
  const [author, setAuthor] = useState('')
  const [content, setContent] = useState('')
  const navigate = useNavigate()

  const handleUserLogin = async (e) => {
    e.preventDefault()
    try {
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          title: title,
          author: author,
          content: content,
        }),
      }
      await fetch('http://localhost:8080/blog', requestOptions)
      navigate("/")
    } catch (err) {
      console.log(err)
    }
  }

  return (
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
  )
}

export default CreateBlog
