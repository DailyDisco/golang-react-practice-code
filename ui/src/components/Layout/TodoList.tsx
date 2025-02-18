/* eslint-disable @typescript-eslint/no-unused-vars */
import { Box, Container, Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { motion } from "framer-motion";
import TodoItem from "./TodoItem";
import { useQuery } from "@tanstack/react-query";
import { BASE_URL } from "../../App"

export type Todo = {
    _id: number;
    body: string;
    completed: boolean;
};

const MotionBox = motion(Box);
const MotionStack = motion(Stack);

const TodoList = () => {
    const { data: todos, isLoading } = useQuery<Todo[]>({
        queryKey: ["todos"],
        queryFn: async () => {
            try {
                const res = await fetch(BASE_URL + "/todos");
                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }
                return data || [];
            } catch (error) {
                console.log(error);
            }
        },
    });

    return (
        <Container maxW="container.md" py={8}>
            <MotionBox
                borderRadius="xl"
                bg="white"
                boxShadow="lg"
                p={6}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5 }}
                border="1px"
                borderColor="gray.300"
                _dark={{
                    bg: "gray.800",
                    borderColor: "gray.700"
                }}
            >
                <Text
                    fontSize={{ base: "3xl", md: "4xl" }}
                    textTransform="uppercase"
                    fontWeight="extrabold"
                    textAlign="center"
                    mb={6}
                    color="blue.700"
                    _dark={{ color: "blue.200" }}
                    letterSpacing="wider"
                >
                    Today's Tasks
                </Text>
                {isLoading && (
                    <Flex justifyContent="center" my={8}>
                        <Spinner
                            size="xl"
                            color="blue.700"
                            _dark={{ color: "blue.200" }}
                        />
                    </Flex>
                )}
                {!isLoading && todos?.length === 0 && (
                    <MotionStack
                        alignItems="center"
                        gap={4}
                        py={8}
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.5 }}
                    >
                        <Text
                            fontSize="xl"
                            textAlign="center"
                            color="gray.700"
                            _dark={{ color: "gray.400" }}
                            fontWeight="medium"
                        >
                            All tasks completed! 🎉
                        </Text>
                        <Box
                            transform="scale(1)"
                            transition="transform 0.2s"
                            _hover={{ transform: "scale(1.1)" }}
                        >
                            <img src="/go.png" alt="Go logo" width={80} height={80} />
                        </Box>
                    </MotionStack>
                )}
                <Stack gap={4} mt={4}>
                    {todos?.map((todo, idx) => (
                        <MotionBox
                            key={idx}
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.5, delay: idx * 0.1 }}
                            borderWidth="1px"
                            borderRadius="lg"
                            borderColor="gray.300"
                            p={2}
                            bg="gray.50"
                            boxShadow="sm"
                            _dark={{
                                borderColor: "gray.700",
                                bg: "gray.800"
                            }}
                        >
                            <TodoItem todo={todo} />
                        </MotionBox>
                    ))}
                </Stack>
            </MotionBox>
        </Container>
    );
};
export default TodoList;

// STARTER CODE:

// import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
// import { useState } from "react";
// import TodoItem from "./TodoItem";

// const TodoList = () => {
//     const [isLoading, setIsLoading] = useState(true);
//     console.log(setIsLoading);

//     const todos = [
//         {
//             _id: 1,
//             body: "Buy groceries",
//             completed: true,
//         },
//         {
//             _id: 2,
//             body: "Walk the dog",
//             completed: false,
//         },
//         {
//             _id: 3,
//             body: "Do laundry",
//             completed: false,
//         },
//         {
//             _id: 4,
//             body: "Cook dinner",
//             completed: true,
//         },
//     ];

//     return (
//         <>
//             <Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
//                 Today's Tasks
//             </Text>
//             {isLoading && (
//                 <Flex justifyContent={"center"} my={4}>
//                     <Spinner size={"xl"} />
//                 </Flex>
//             )}
//             {!isLoading && todos?.length === 0 && (
//                 <Stack alignItems={"center"} gap='3'>
//                     <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
//                         All tasks completed! 🤞
//                     </Text>
//                     <img src='/go.png' alt='Go logo' width={70} height={70} />
//                 </Stack>
//             )}
//             <Stack gap={3}>
//                 {todos?.map((todo) => (
//                     <TodoItem key={todo._id} todo={todo} />
//                 ))}
//             </Stack>
//         </>
//     );
// };

// export default TodoList;