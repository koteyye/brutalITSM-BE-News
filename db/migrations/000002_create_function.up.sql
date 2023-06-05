
create function getNewsList()
    returns table (id uuid, title text, description text, created_at timestamp, user_created uuid, updated_at timestamp, user_updated uuid, state news_state, previewImage json, contentFile json)
    language plpgsql
as
$$
begin
    return query
        select n.id, n.title, n.description, n.created_at, n.user_created, n.updated_at, n.user_updated, n.state,
               json_build_object(
                       'bucketName', pf.bucket_name,
                       'fileName', pf.file_name,
                       'mimeType', pf.mime_type
                   ) previewImage,
               json_build_object(
                       'bucketName', pc.bucket_name,
                       'fileName', pc.file_name,
                       'mimeType', pc.mime_type
                   ) contentFile
        from news n
                 left join files pf on pf.id = n.preview_image
                 left join files pc on pc.id = n.content_file
        where n.deleted_at is null and n.state = 'published';
end;
$$;

create function getNewsById(userId uuid)
    returns table (id uuid, title text, description text, created_at timestamp, user_created uuid, updated_at timestamp, user_updated uuid, state news_state, previewImage json, contentFile json)
    language plpgsql
as
$$
begin
    return query
        select n.id, n.title, n.description, n.created_at, n.user_created, n.updated_at, n.user_updated, n.state,
               json_build_object(
                       'bucketName', pf.bucket_name,
                       'fileName', pf.file_name,
                       'mimeType', pf.mime_type
                   ) previewImage,
               json_build_object(
                       'bucketName', pc.bucket_name,
                       'fileName', pc.file_name,
                       'mimeType', pc.mime_type
                   ) contentFile
        from news n
                 left join files pf on pf.id = n.preview_image
                 left join files pc on pc.id = n.content_file
        where n.id = userId and n.deleted_at is null and n.state = 'published';
end;
$$;

