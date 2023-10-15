import {
  Alert,
  AlertIcon,
  AlertTitle,
  Box,
  Button,
  Image,
  Input,
  Tab,
  TabList,
  Tabs
} from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';


const placeholderImage = 'https://via.placeholder.com/512';

export default function UploadForm() {
  const [file, setFile] = useState(null);
  const [selectedModelId, setSelectedModelId] = useState(''); // 選択したモデルのIDを管理
  const [selectedModelName, setSelectedModelName] = useState(''); // 選択したモデルの名前を管理
  const [responseMessage, setResponseMessage] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [loading, setLoading] = useState(false);
  const [models, setModels] = useState([]);
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
        console.log(response.data);
      } catch (error) {
        console.error(error);
        // エラーが発生した場合に/loginにリダイレクト
        router.push('/login');
      }
    }

    const fetchData = async () => {
      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/models`, {
          headers: {
            'Authorization': 'Bearer ' + token,
          },
        });
        const modelData = response.data.prediction_models;
        setModels(modelData);

        if (modelData.length > 0) {
          setSelectedModelId(modelData[0].id); // 最初のモデルのIDをデフォルトで選択
          setSelectedModelName(modelData[0].model_name); // 最初のモデルの名前をデフォルトで選択
        }
      } catch (error) {
        console.error(error);
      }
    };

    fetchUser();
    fetchData();
  }, [token]);

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    setFile(selectedFile);

    const reader = new FileReader();
    reader.onload = (e) => {
      setImagePreview(e.target.result);
    };
    reader.readAsDataURL(selectedFile);
  };

  const handleTabChange = (model) => {
    setSelectedModelId(model.id); // タブが変更されたら選択したモデルのIDを更新
    setSelectedModelName(model.model_name); // タブが変更されたら選択したモデルの名前を更新
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append('image', file);

    try {
      setLoading(true);
      const response = await axios.post(`${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/predict/${selectedModelId}`, formData, {
        headers: {
          'Authorization': 'Bearer ' + token,
        },
      });
      const responseData = response.data;
      setResponseMessage(responseData.prediction_result.result);
    } catch (error) {
      console.error(error);
      setResponseMessage('エラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box maxW="fit-content">
      <form onSubmit={handleSubmit}>
        {responseMessage && (
        <Alert status={responseMessage != 'エラーが発生しました' ? 'success' : 'error'}>
          <AlertIcon />
          <AlertTitle>{responseMessage}</AlertTitle>
        </Alert>
        )}
        <Box paddingLeft={4} marginTop={4}>
          {imagePreview ? (
            <Image src={imagePreview} alt="プレビュー" boxSize="224" borderRadius="10%" />
          ) : (
            <Image src={placeholderImage} alt="プレースホルダー" boxSize="224" borderRadius="10%" />
          )}
          <Box>
            <Input type="file" onChange={handleFileChange} required accept=".png, .jpg, .jpeg" />
          </Box>
          <Box marginTop={4}>
            <Tabs isFitted variant="solid-rounded" colorScheme="cyan">
              <TabList>
                {models.map((model) => (
                  <Tab key={model.id} onClick={() => handleTabChange(model)}>
                    {model.model_name}
                  </Tab>
                ))}
              </TabList>
            </Tabs>
          </Box>
          <Box marginTop={4}>
            <Button type="submit" colorScheme='teal' variant='solid' size="lg" onClick={handleSubmit} isLoading={loading} loadingText="診断中">診断開始</Button>
          </Box>
        </Box>
      </form>
    </Box>
  );
}
