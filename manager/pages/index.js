import UploadForm from "@/components/UploadForm";
import { Box } from "@chakra-ui/react";
import Header from "../components/Header";
import useAuth from "../hooks/useAuth";

export default function Home() {
  useAuth();
  return (
    <Box>
      <Header />
      <UploadForm />
    </Box>
  );
};
