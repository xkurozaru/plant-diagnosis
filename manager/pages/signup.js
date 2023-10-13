import Header from "@/components/Header";
import { Alert, AlertIcon, Box, Heading } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";
import AuthForm from "../components/AuthForm";

const SignupPage = () => {
  const [errorAlert, setErrorAlert] = useState(false);

  const handleSignup = async (userData) => {
    setErrorAlert(false); // リクエストを送信する前にアラートを隠す
    // サインアップ処理を実装する
    try {
      const response = await axios.post(`${process.env.NEXT_PUBLIC_HOST}/api/v1/sign-up`, userData);
      console.log("サインアップ成功:", response.data);
      // サインアップが成功した場合の処理を追加
      window.location.href = "/login";

    } catch (error) {
      console.log("サインアップエラー:", error);
      // サインアップが失敗した場合のエラー処理を追加
      setErrorAlert(true); // エラーアラートを表示
    }
  };

  return (
    <Box>
      <Header />
      {errorAlert && (
        <Alert status="error">
          <AlertIcon />
          そのログインIDは既に使われています
        </Alert>
      )}
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
