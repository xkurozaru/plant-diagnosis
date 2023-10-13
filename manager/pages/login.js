import { Alert, AlertIcon, Box, Heading } from "@chakra-ui/react";
import axios from "axios";
import Cookies from "js-cookie";
import { useState } from "react";
import AuthForm from "../components/AuthForm";
import Header from "../components/Header";


const LoginPage = () => {
  const [errorAlert, setErrorAlert] = useState(false);
  const handleLogin = async (userData) => {
    setErrorAlert(false); // リクエストを送信する前にアラートを隠す
    try {
      const response = await axios.post(`${process.env.NEXT_PUBLIC_HOST}/api/v1/sign-in`, userData);
      const token  = response.data.token; // レスポンスからJWTを取得

      // JWTをCookieに保存
      Cookies.set("token", token);

      // ログインが成功した場合の処理を追加
      window.location.href = "/";

    } catch (error) {
      // ログインが失敗した場合のエラー処理を追加
      setErrorAlert(true); // エラーアラートを表示
    }
  };

  return (
    <Box>
      <Header />
      {errorAlert && (
        <Alert status="error">
          <AlertIcon />
          ログインIDまたはパスワードが間違っています
        </Alert>
      )}
      <Box paddingLeft={10} marginTop={4}>
        <Heading as="h2" size="xl" marginBottom={4}>
          ログイン
        </Heading>
        <AuthForm onSubmit={handleLogin} buttonText="ログイン" hideUsernameField={true} />
      </Box>
    </Box>
  );
};

export default LoginPage;
