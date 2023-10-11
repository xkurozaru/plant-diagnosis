import Header from '@/components/Header';
import { Box, Flex } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import ModelForm from '../components/ModelForm';
import ModelList from '../components/ModelList';


const AdminPage = () => {
  const token = Cookies.get('token');
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get('http://localhost:8000/api/v1/users', {
          headers: {
            'Authorization': 'Bearer ' + token,
          },
        });
        console.log(response.data.user.role);
        if (response.data.user.role !== 'admin') {
          router.push('/login'); // 管理者以外は/loginにリダイレクト
        }
      } catch (error) {
        console.error(error);
        router.push('/login'); // エラーが発生した場合に/loginにリダイレクト
      }
    }

    fetchUser();
  }, [token]);

  return (
    <Box>
      <Header />
      <Flex marginTop={4} paddingLeft={4} gap={2}>
        <ModelForm />
        <ModelList />
      </Flex>
    </Box>
  );
};

export default AdminPage;
