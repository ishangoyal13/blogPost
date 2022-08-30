import React, { useState } from 'react'
import Navbar from 'react-bootstrap/Navbar'
import Container from 'react-bootstrap/Container'
import Button from 'react-bootstrap/Button'
import { FiEdit } from 'react-icons/fi'
import {TiArrowBackOutline} from 'react-icons/ti'
import { useNavigate } from 'react-router-dom'

const Header = () => {
  const navigate = useNavigate()
  const [onEditPage, setOnEditPage] = useState(false)

  return (
    <>
      <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand>Welcome User :)</Navbar.Brand>
          <Navbar.Toggle />
          <Navbar.Collapse className="justify-content-end">
            <Navbar.Text>
              {!onEditPage ? (
                <Button
                  variant="primary"
                  onClick={() => {
                    setOnEditPage(true)
                    navigate('/addblog')
                  }}
                >
                  Add Blog <FiEdit />
                </Button>
              ) : (
                <Button
                  variant="primary"
                  onClick={() => {
                    setOnEditPage(false)
                    navigate('/home')
                  }}
                >
                  Go Back <TiArrowBackOutline />
                </Button>
              )}
            </Navbar.Text>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  )
}

export default Header
