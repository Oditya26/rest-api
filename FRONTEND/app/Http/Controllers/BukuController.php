<?php

namespace App\Http\Controllers;

use GuzzleHttp\Client;
use Illuminate\Http\Request;

class BukuController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $client = new Client();
        $url = "http://127.0.0.1:8080/api/buku";
        $response = $client->request('GET', $url);
        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        $data = $contentArray['data'];
        return view('buku.index',['data'=>$data]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $currentTime = date('Y-m-d H:i:s');
        $judul = $request->judul;
        $pengarang = $request->pengarang;
        $tanggal_publikasi = $request->tanggal_publikasi;
        $created_at =  $currentTime;
        $updated_at =  null;

        $parameter = [
            'judul'=>$judul,
            'pengarang'=>$pengarang,
            'tanggal_publikasi'=>$tanggal_publikasi,
            'created_at'=>$created_at,
            'updated_at'=>$updated_at,
        ];

        $client = new Client();
        $url = "http://127.0.0.1:8080/api/buku";
        $response = $client->request('POST', $url,[
            'headers'=>['Content-type'=>'application/json'],
            'body'=>json_encode($parameter),
        ]);
        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        
        if($contentArray['status']!=true) {
            $error = $contentArray['data'];
            return redirect()->to('buku')->withErrors($error)->withInput();
        } else
        {
            return redirect()->to('buku')->with('success','Berhasil memasukkan data.');
        }
        
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(string $id)
    {
        $client = new Client();
        $url = "http://127.0.0.1:8080/api/buku/$id";
        $response = $client->request('GET', $url);
        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        if($contentArray['status']!=true) {
            $error = $contentArray['message'];
            return redirect()->to('buku')->withErrors($error);
        } else {
            $data = $contentArray['data'];
            return view('buku.index', ['data'=>$data]);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        $currentTime = date('Y-m-d H:i:s');
        $judul = $request->judul;
        $pengarang = $request->pengarang;
        $tanggal_publikasi = $request->tanggal_publikasi;
        $updated_at =  $currentTime;

        $parameter = [
            'judul'=>$judul,
            'pengarang'=>$pengarang,
            'tanggal_publikasi'=>$tanggal_publikasi,
            'updated_at'=>$updated_at,
        ];

        $client = new Client();
        $url = "http://127.0.0.1:8080/api/buku/$id";
        $response = $client->request('PUT', $url,[
            'headers'=>['Content-type'=>'application/json'],
            'body'=>json_encode($parameter),
        ]);
        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        if($contentArray['status']!=true) {
            $error = $contentArray['data'];
            return redirect()->to('buku')->withErrors($error)->withInput();
        } else
        {
            return redirect()->to('buku')->with('success','Berhasil melakukan update data.');
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $client = new Client();
        $url = "http://127.0.0.1:8080/api/buku/$id";
        $response = $client->request('DELETE', $url);
        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        if($contentArray['status']!=true) {
            $error = $contentArray['data'];
            return redirect()->to('buku')->withErrors($error)->withInput();
        } else
        {
            return redirect()->to('buku')->with('success','Berhasil melakukan delete data.');
        }
    }
}
