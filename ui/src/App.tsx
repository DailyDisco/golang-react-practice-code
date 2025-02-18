// /* eslint-disable @typescript-eslint/no-unused-vars */
// import { useState } from 'react'

import { Stack, Container } from '@chakra-ui/react'

import Navbar from './components/Layout/Navbar'
import TodoForm from './components/Layout/TodoForm'
import TodoList from './components/Layout/TodoList'
import FilterTodos from './components/Layout/FilterTodos'

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";

function App() {
  // State Example:
  // const [count, setCount] = useState(0)
  // console.log(count, "count")

  return (
    <Stack h="100vh">
      <Navbar />
      <Container maxW="3/5">
        <TodoForm />
        <FilterTodos />
        <TodoList />
      </Container>
    </Stack>
  )
}

export default App
