import { Box, Button, FormControl, FormLabel, Input } from "@chakra-ui/react";

const AuthForm = ({ onSubmit, buttonText, hideUsernameField = false }) => {
  const handleSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const userData = {
      login_id: formData.get("login_id"),
      password: formData.get("password"),
    };
    if (!hideUsernameField) {
      userData.user_name = formData.get("user_name");
    }
    onSubmit(userData);
  };

  return (
    <Box maxW="fit-content">
      <form onSubmit={handleSubmit}>
        {!hideUsernameField && (
          <FormControl>
            <FormLabel>ユーザー名</FormLabel>
            <Input type="text" name="user_name" required />
          </FormControl>
        )}
        <FormControl mt={4}>
          <FormLabel>ログインID</FormLabel>
          <Input type="text" name="login_id" required minLength={4} maxLength={32} />
        </FormControl>
        <FormControl mt={4}>
          <FormLabel>パスワード</FormLabel>
          <Input type="password" name="password" required minLength={8} maxLength={72} />
        </FormControl>
        <Button colorScheme="teal" mt={4} type="submit">
          {buttonText}
        </Button>
      </form>
    </Box>
  );
};

export default AuthForm;
