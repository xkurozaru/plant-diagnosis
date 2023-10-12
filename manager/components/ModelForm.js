import { Box, Button, Input, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useState } from 'react';

const ModelForm = () => {
  const [modelData, setModelData] = useState({
    model_name: '',
    network_name: '',
    param_path: '',
    labels: [],
  });
  const [loading, setLoading] = useState(false);
  const token = Cookies.get('token');

  const sendRequest = async () => {
    setLoading(true);
    const url = 'http://localhost:8000/api/v1/prediction/models';
    try {
      const response = await axios.post(url, modelData, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 200) {
        // リクエストが成功した場合の処理
        console.log('リクエストが成功しました');
        // ここで必要な処理を追加できます
      } else {
        // リクエストが失敗した場合の処理
        console.error('リクエストが失敗しました');
        // エラーハンドリングを追加できます
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setModelData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  return (
    <VStack spacing={4}>
      <Text fontSize="xl" fontWeight="bold">
        モデル作成
      </Text>
      <Box>
        <Text>モデル名:</Text>
        <Input
          type="text"
          name="model_name"
          value={modelData.model_name}
          onChange={handleChange}
        />
      </Box>
      <Box>
        <Text>ネットワーク名:</Text>
        <Input
          type="text"
          name="network_name"
          value={modelData.network_name}
          onChange={handleChange}
        />
      </Box>
      <Box>
        <Text>パラメーターパス:</Text>
        <Input
          type="text"
          name="param_path"
          value={modelData.param_path}
          onChange={handleChange}
        />
      </Box>
      <Button colorScheme='teal' variant='solid' size="lg" onClick={sendRequest} isLoading={loading} loadingText="モデル作成中">モデル作成</Button>
    </VStack>
  );
};

export default ModelForm;
