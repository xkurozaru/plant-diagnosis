import Header from "@/components/header";
import { Container, Heading } from "@chakra-ui/react";
import { Router } from "next/router";
import AuthForm from "../components/AuthForm";

const SignupPage = () => {
  const handleSignup = async (userData) => {
    // サインアップ処理を実装する
    try {
      const response = await axios.post(`http://localhost:8000/api/sign-up`, userData);
      console.log("サインアップ成功:", response.data);
      // サインアップが成功した場合の処理を追加
      Router.push("/login"); // サインアップ後にログインページに遷移

    } catch (error) {
      console.log("サインアップエラー:", error);
      // サインアップが失敗した場合のエラー処理を追加
    }
  };

  return (
    <Container maxW="sm" centerContent>
      <Header />
      <Heading as="h2" size="xl" mt={8}>
        サインアップ
      </Heading>
      <AuthForm onSubmit={handleSignup} buttonText="サインアップ" />
    </Container>
  );
};

export default SignupPage;
