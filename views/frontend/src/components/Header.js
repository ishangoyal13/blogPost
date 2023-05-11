import React from 'react'
import Navbar from 'react-bootstrap/Navbar'
import Container from 'react-bootstrap/Container'
import Button from 'react-bootstrap/Button'
import { FiEdit,FiLogOut } from 'react-icons/fi'
import { TiArrowBackOutline } from 'react-icons/ti'
import { useNavigate, useLocation } from 'react-router-dom'

const Header = () => {
  const navigate = useNavigate()
  const state = useLocation()
  return (
    <>
      <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand>Welcome {state.state.first_name} :)</Navbar.Brand>
          <Navbar.Toggle />
          <Navbar.Collapse className="justify-content-end">
            <Navbar.Text>
              {state.pathname !== "/addblog" ? (
                <Button
                  variant="primary"
                  onClick={() => {
                    navigate('/addblog', { state: { first_name: state.state.first_name } })
                  }}
                >
                  Add Blog <FiEdit />
                </Button>
              ) : (
                <Button
                  variant="primary"
                  onClick={() => {
                    navigate('/home', { state: { first_name: state.state.first_name } })
                  }}
                >
                  Go Back <TiArrowBackOutline />
                </Button>
              )}
            </Navbar.Text>&nbsp;
            <Navbar.Text>
              <Button
                variant="primary"
                onClick={() => {
                  navigate('/login')
                }}
              >
                Log Out <FiLogOut />
              </Button>
            </Navbar.Text>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  )
}

export default Header
