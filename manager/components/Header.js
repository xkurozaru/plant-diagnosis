import { Box, Flex, Heading, Spacer } from '@chakra-ui/react';
import axios from "axios";
import Cookies from "js-cookie";
import NextLink from "next/link";
import { useEffect, useState } from "react";

export default function Header() {
  const [userName, setUserName] = useState(null);

  useEffect(() => {
    // CookieからTokenを取得
    const token = Cookies.get("token");

    if (token) {
      // Tokenがある場合はAPIを呼び出してユーザー名を取得
      axios.get(`${process.env.NEXT_PUBLIC_HOST}/api/v1/users`, {
        headers: {
          'Authorization': 'Bearer ' + token, // 認証トークン
          'Content-Type': 'application/json', // コンテンツタイプも指定できます
        },
      })
      .then((response) => {
        setUserName(response.data.user.name);
      })
      .catch((error) => {
        console.error("ユーザー名の取得に失敗しました", error);
      });
    }
  }, []);

  return (
    <Box as="header" bg="gray.100" px="6" py="6">
      <Flex alignItems="center">
        <NextLink href="/">
          <Heading as="h1" size="lg" cursor="pointer">
            植物病害診断
          </Heading>
        </NextLink>
        <Spacer />
        <Box>
          {userName ? (
            // Tokenが保存されている場合、ユーザー名を表示
            <p>ようこそ、{userName}さん</p>
          ) : (
            // Tokenが保存されていない場合、signupへのリンクを表示
            <NextLink href="/signup">
              <p>新規登録</p>
            </NextLink>
          )}
        </Box>
      </Flex>
    </Box>
  );
}
