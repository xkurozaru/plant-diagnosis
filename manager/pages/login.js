import { Container, Heading } from "@chakra-ui/react";
import axios from "axios";
import Cookies from "js-cookie"; // js-cookieをインポート
import { Router } from "next/router";
import AuthForm from "../components/AuthForm";

const LoginPage = () => {
  const handleLogin = async (userData) => {
    try {
      const response = await axios.post("http://localhost:8000/api/sign-in", userData);
      const { token } = response.data; // レスポンスからJWTを取得

      // JWTをCookieに保存
      Cookies.set("token", token);

      console.log("ログイン成功:", response.data);
      // ログインが成功した場合の処理を追加
      Router.push("/"); // ログイン後にトップページに遷移

    } catch (error) {
      console.error("ログインエラー:", error.response.data);
      // ログインが失敗した場合のエラー処理を追加
    }
  };

  return (
    <Container maxW="sm" centerContent>
      <Heading as="h2" size="xl" mt={8}>
        ログイン
      </Heading>
      <AuthForm onSubmit={handleLogin} buttonText="ログイン" hideUsernameField={true} />
    </Container>
  );
};

export default LoginPage;
