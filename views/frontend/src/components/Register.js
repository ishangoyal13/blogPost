import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

function Register() {
  const [name, setName] = useState('')
  const [phoneNumber, setPhoneNumber] = useState()
  const [pass, setPass] = useState('')
  const navigate = useNavigate()

  const handleUserLogin = async e => {
    e.preventDefault()
    axios.post("http://localhost:8000/api/user/register", {
      name: name,
      phone_number: parseInt(phoneNumber),
      password: pass,
    }).then(response => {
      if (response.status === 201) {
        localStorage.setItem("token", response.data.token)
        navigate('/home', { state: { first_name: name } })
      }
    }).catch(err => {
      console.log(err.response.data)
    })
  }

  const ShowPassword = () => {
    var x = document.getElementById('password')
    if (x.type === "password") {
      x.type = "text"
    } else {
      x.type = "password"
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

        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Phone Number</Form.Label>
          <Form.Control
            type="text"
            placeholder="Enter Phone number"
            value={phoneNumber}
            onChange={e => {
              setPhoneNumber(e.target.value)
            }}
          />
        </Form.Group>

        <Form.Group className="mb-3">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            value={pass}
            onChange={e => {
              setPass(e.target.value)
            }}
            id='password'
          />
          <Form.Switch
            label="Show Password"
            onClick={ShowPassword}
            style={{ float: 'right' }}
          >
          </Form.Switch>
        </Form.Group>

        <Form.Group className="mb-3">
          <br/>
          <Button
            variant="primary"
            type="submit"
            onClick={e => handleUserLogin(e)}
          >
            Register User
          </Button>
          <p style={{ float: 'right' }}>Already a User ? <a href='/login' style={{}}>LogIN</a></p>
        </Form.Group>
      </Form>
    </div>
  )
}

export default Register
