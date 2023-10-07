import Header from "@/components/Header";
import { Box, Heading } from "@chakra-ui/react";
import axios from "axios";
import AuthForm from "../components/AuthForm";

const SignupPage = () => {
  const handleSignup = async (userData) => {
    // サインアップ処理を実装する
    try {
      const response = await axios.post(`http://localhost:8000/api/v1/sign-up`, userData);
      console.log("サインアップ成功:", response.data);
      // サインアップが成功した場合の処理を追加
      window.location.href = "/login";

    } catch (error) {
      console.log("サインアップエラー:", error);
      // サインアップが失敗した場合のエラー処理を追加
    }
  };

  return (
    <Box>
      <Header />
      <Box paddingLeft={10} marginTop={4}>
        <Heading as="h2" size="xl" marginBottom={4}>
          サインアップ
        </Heading>
        <AuthForm onSubmit={handleSignup} buttonText="サインアップ" />
      </Box>
    </Box>
  );
};

export default SignupPage;
