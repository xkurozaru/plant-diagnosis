import { Box, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

const ModelList = () => {
  const [models, setModels] = useState([]);
  const [loading, setLoading] = useState(true);
  const token = Cookies.get('token');
  const router = useRouter();

  useEffect(() => {
    const fetchModels = async () => {
      try {
        const url = 'http://localhost:8000/api/v1/prediction/models';
        const response = await axios.get(url, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.status === 200) {
          // リクエストが成功した場合の処理
          setModels(response.data.prediction_models);
          setLoading(false);
        } else {
          // リクエストが失敗した場合の処理
          console.error('リクエストが失敗しました');
          setLoading(false);
          // エラーハンドリングを追加できます
        }
      } catch (error) {
        console.error('エラーが発生しました:', error);
        setLoading(false);
      }
    };

    fetchModels();
  }, []);

  return (
    <VStack spacing={4}>
      <Text fontSize="xl" fontWeight="bold">
        モデル一覧
      </Text>
      {loading ? (
        <Text>データを読み込んでいます...</Text>
      ) : (
        models.map((model) => (
          <Box key={model.id} borderWidth="1px" p={4} borderRadius="md">
            <Text fontSize="lg">{model.model_name}</Text>
            <Text>ラベル: {model.labels.join(', ')}</Text>
          </Box>
        ))
      )}
    </VStack>
  );
};

export default ModelList;
