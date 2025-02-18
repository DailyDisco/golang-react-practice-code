import { Button, ButtonGroup, Text, VStack } from "@chakra-ui/react"
import { useQueryClient } from "@tanstack/react-query"
import { useState } from "react"
import { Todo } from "./TodoList"

type FilterType = "all" | "completed" | "in-progress"

const FilterTodos = () => {
    const [activeFilter, setActiveFilter] = useState<FilterType>("all")
    const queryClient = useQueryClient()

    const handleFilter = (filter: FilterType) => {
        setActiveFilter(filter)
        // Always get the original data
        const originalTodos = queryClient.getQueryData<Todo[]>(["todos-original"])

        if (!originalTodos) {
            // If no original data exists, get and store it
            const currentTodos = queryClient.getQueryData<Todo[]>(["todos"]) || []
            queryClient.setQueryData(["todos-original"], currentTodos)
            return
        }

        let filteredTodos
        switch (filter) {
            case "completed":
                filteredTodos = originalTodos.filter(todo => todo.completed)
                break
            case "in-progress":
                filteredTodos = originalTodos.filter(todo => !todo.completed)
                break
            default:
                filteredTodos = originalTodos
        }

        queryClient.setQueryData(["todos"], filteredTodos)
    }

    return (
        <VStack gap={4} my={4}>
            <Text
                fontSize="lg"
                fontWeight="medium"
                color="gray.700"
                _dark={{ color: "gray.200" }}
            >
                Filter Tasks
            </Text>
            <ButtonGroup variant="outline">
                <Button
                    size="sm"
                    variant="outline"
                    colorScheme={activeFilter === "all" ? "blue" : "gray"}
                    onClick={() => handleFilter("all")}
                >
                    All
                </Button>
                <Button
                    size="sm"
                    variant="outline"
                    colorScheme={activeFilter === "completed" ? "green" : "gray"}
                    onClick={() => handleFilter("completed")}
                >
                    Completed
                </Button>
                <Button
                    size="sm"
                    variant="outline"
                    colorScheme={activeFilter === "in-progress" ? "yellow" : "gray"}
                    onClick={() => handleFilter("in-progress")}
                >
                    In Progress
                </Button>
            </ButtonGroup>
        </VStack>
    )
}

export default FilterTodos