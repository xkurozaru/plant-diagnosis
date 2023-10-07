import Cookies from 'js-cookie';
import { useEffect } from 'react';

export default function useAuth() {
  useEffect(() => {
    // Cookieからトークンを取得
    const token = Cookies.get('token');

    // ログイン状態の確認
    if (!token) {
      // 未ログインの場合、ログインページにリダイレクト
      router.push('/login');
    }
  }, []);
}
