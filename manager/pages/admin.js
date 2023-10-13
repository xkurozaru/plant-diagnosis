import Header from '@/components/Header';
import { Box, Flex } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import ModelForm from '../components/ModelForm';
import ModelList from '../components/ModelList';


const AdminPage = () => {
  const [models, setModels] = useState([]); // モデルリストの状態を管理
  const token = Cookies.get('token');
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_HOST}/api/v1/users`, {
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
    updateModelList(); // 初回読み込み時にモデルリストを取得
  }, [token]);

  const updateModelList = async () => {
    try {
      const url = `${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/models`;
      const response = await axios.get(url, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 200) {
        setModels(response.data.prediction_models); // モデルリストを更新
      } else {
        console.error('リクエストが失敗しました');
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
    }
  };

  return (
    <Box>
      <Header />
      <Flex marginTop={4} paddingLeft={4} gap={2}>
        <ModelForm updateModelList={updateModelList} />
        <ModelList models={models} setModels={setModels} />
      </Flex>
    </Box>
  );
};

export default AdminPage;
