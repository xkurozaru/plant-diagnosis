import {
  Box,
  Heading
} from '@chakra-ui/react';
import NextLink from "next/link";

export default function Header() {
  return (
    <Box as="header" bg="gray.100" px="6" py="6" mb="6">
      <NextLink href="/">
        <Heading as="h1" size="lg" cursor="pointer">
          植物病害診断
        </Heading>
      </NextLink>
    </Box>
  );
}
