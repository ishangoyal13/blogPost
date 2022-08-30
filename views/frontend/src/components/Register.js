import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom'

function Register() {
  const [name, setName] = useState('')
  const [username, setUser] = useState('')
  const [email, setEmail] = useState('')
  const [pass, setPass] = useState('')
  const navigate = useNavigate()

  const handleUserLogin = async e => {
    e.preventDefault()
    try {
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          name: name,
          username: username,
          email: email,
          password: pass,
        }),
      }
      await fetch('http://localhost:8080/api/user/register', requestOptions)
      navigate('/login')
    } catch (err) {
      console.log(err)
    }
  }
  return (
    <div className="createDiv">
      <Form>
        <Form.Group className="mb-3">
          <Form.Label>Name</Form.Label>
          <Form.Control
            type="text"
            placeholder="Enter name"
            value={name}
            onChange={e => {
              setName(e.target.value)
            }}
          />
        </Form.Group>

        <Form.Group className="mb-3">
          <Form.Label>Username</Form.Label>
          <Form.Control
            type="text"
            placeholder="Enter username"
            value={username}
            onChange={e => {
              setUser(e.target.value)
            }}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Email</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={e => {
              setEmail(e.target.value)
            }}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            value={pass}
            onChange={e => {
              setPass(e.target.value)
            }}
          />
        </Form.Group>

        <Button
          variant="primary"
          type="submit"
          onClick={e => handleUserLogin(e)}
        >
          Register User
        </Button>
      </Form>
    </div>
  )
}

export default Register
